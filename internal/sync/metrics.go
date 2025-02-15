package sync

import (
	"fmt"
	"github.com/amazechain/amc/internal/p2p"
	"github.com/amazechain/amc/log"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	topicPeerCount = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "p2p_topic_peer_count",
			Help: "The number of peers subscribed to a given topic.",
		}, []string{"topic"},
	)
	subscribedTopicPeerCount = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "p2p_subscribed_topic_peer_total",
			Help: "The number of peers subscribed to topics that a host node is also subscribed to.",
		}, []string{"topic"},
	)
	messageReceivedCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "p2p_message_received_total",
			Help: "Count of messages received.",
		},
		[]string{"topic"},
	)
	messageFailedValidationCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "p2p_message_failed_validation_total",
			Help: "Count of messages that failed validation.",
		},
		[]string{"topic"},
	)
	messageIgnoredValidationCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "p2p_message_ignored_validation_total",
			Help: "Count of messages that were ignored in validation.",
		},
		[]string{"topic"},
	)
	messageFailedProcessingCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "p2p_message_failed_processing_total",
			Help: "Count of messages that passed validation but failed processing.",
		},
		[]string{"topic"},
	)
	numberOfTimesResyncedCounter = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "number_of_times_resynced",
			Help: "Count the number of times a node resyncs.",
		},
	)
	duplicatesRemovedCounter = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "number_of_duplicates_removed",
			Help: "Count the number of times a duplicate signature set has been removed.",
		},
	)
	numberOfSetsAggregated = promauto.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "number_of_sets_aggregated",
			Help:    "Count the number of times different sets have been successfully aggregated in a batch.",
			Buckets: []float64{10, 50, 100, 200, 400, 800, 1600, 3200},
		},
	)
	rpcBlocksByRangeResponseLatency = promauto.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "rpc_blocks_by_range_response_latency_milliseconds",
			Help:    "Captures total time to respond to rpc blocks by range requests in a milliseconds distribution",
			Buckets: []float64{5, 10, 50, 100, 150, 250, 500, 1000, 2000},
		},
	)
	arrivalBlockPropagationHistogram = promauto.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "block_arrival_latency_milliseconds",
			Help:    "Captures blocks propagation time. Blocks arrival in milliseconds distribution",
			Buckets: []float64{100, 250, 500, 750, 1000, 1500, 2000, 4000, 8000, 12000, 16000, 20000, 24000},
		},
	)
	arrivalBlockPropagationGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "block_arrival_latency_milliseconds_gauge",
		Help: "Captures blocks propagation time. Blocks arrival in milliseconds",
	})

	// Attestation processing granular error tracking.
	attBadBlockCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: "gossip_attestation_bad_block_total",
		Help: "Increased when a gossip attestation references a bad block",
	})
	attBadLmdConsistencyCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: "gossip_attestation_bad_lmd_consistency_total",
		Help: "Increased when a gossip attestation has bad LMD GHOST consistency",
	})
	attBadSelectionProofCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: "gossip_attestation_bad_selection_proof_total",
		Help: "Increased when a gossip attestation has a bad selection proof",
	})
	attBadSignatureBatchCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: "gossip_attestation_bad_signature_batch_total",
		Help: "Increased when a gossip attestation has a bad signature batch",
	})

	// Attestation and block gossip verification performance.
	aggregateAttestationVerificationGossipSummary = promauto.NewSummary(
		prometheus.SummaryOpts{
			Name: "gossip_aggregate_attestation_verification_milliseconds",
			Help: "Time to verify gossiped attestations",
		},
	)
	blockVerificationGossipSummary = promauto.NewSummary(
		prometheus.SummaryOpts{
			Name: "gossip_block_verification_milliseconds",
			Help: "Time to verify gossiped blocks",
		},
	)
)

func (s *Service) updateMetrics() {
	// do not update metrics if genesis time
	// has not been initialized
	if s.cfg.chain.CurrentBlock().Header().Number64().IsZero() {
		return
	}
	// We update the dynamic subnet topics.
	digest, err := s.currentForkDigest()
	if err != nil {
		log.Debugf("Could not compute fork digest", "err", err)
	}

	// We update all other gossip topics.
	for _, topic := range p2p.AllTopics() {
		topic += s.cfg.p2p.Encoding().ProtocolSuffix()
		if !strings.Contains(topic, "%x") {
			topicPeerCount.WithLabelValues(topic).Set(float64(len(s.cfg.p2p.PubSub().ListPeers(topic))))
			continue
		}
		formattedTopic := fmt.Sprintf(topic, digest)
		topicPeerCount.WithLabelValues(formattedTopic).Set(float64(len(s.cfg.p2p.PubSub().ListPeers(formattedTopic))))
	}

	for _, topic := range s.cfg.p2p.PubSub().GetTopics() {
		subscribedTopicPeerCount.WithLabelValues(topic).Set(float64(len(s.cfg.p2p.PubSub().ListPeers(topic))))
	}
}
