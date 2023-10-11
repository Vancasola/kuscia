package metric_types

// NewMetricTypes parse the metric types from a yaml file
func NewMetricTypes() map[string]string {
	MetricTypes := make(map[string]string)
	MetricTypes["rto"] = "Gauge"
	MetricTypes["rtt"] = "Gauge"
	MetricTypes["delivery_rate"] = "Gauge"
	MetricTypes["bytes_send"] = "Counter"
	MetricTypes["bytes_received"] = "Counter"
	MetricTypes["min_rtt"] = "Gauge"
	MetricTypes["assignment_stale"] = "Counter"
	MetricTypes["assignment_timeout_received"] = "Counter"
	MetricTypes["bind_errors"] = "Counter"
	MetricTypes["circuit_breakers__default__cx_open"] = "Gauge"
	MetricTypes["circuit_breakers__default__cx_pool_open"] = "Gauge"
	MetricTypes["circuit_breakers__default__rq_open"] = "Gauge"
	MetricTypes["circuit_breakers__default__rq_pending_open"] = "Gauge"
	MetricTypes["circuit_breakers__default__rq_retry_open"] = "Gauge"
	MetricTypes["circuit_breakers__high__cx_open"] = "Gauge"
	MetricTypes["circuit_breakers__high__cx_pool_open"] = "Gauge"
	MetricTypes["circuit_breakers__high__rq_open"] = "Gauge"
	MetricTypes["circuit_breakers__high__rq_pending_open"] = "Gauge"
	MetricTypes["circuit_breakers__high__rq_retry_open"] = "Gauge"
	MetricTypes["default__total_match_count"] = "Counter"
	MetricTypes["health_check__attempt"] = "Counter"
	MetricTypes["health_check__degraded"] = "Counter"
	MetricTypes["health_check__failure"] = "Counter"
	MetricTypes["health_check__healthy"] = "Gauge"
	MetricTypes["health_check__network_failure"] = "Counter"
	MetricTypes["health_check__passive_failure"] = "Counter"
	MetricTypes["health_check__success"] = "Counter"
	MetricTypes["health_check__verify_cluster"] = "Counter"
	MetricTypes["http1__dropped_headers_with_underscores"] = "Counter"
	MetricTypes["http1__metadata_not_supported_error"] = "Counter"
	MetricTypes["http1__requests_rejected_with_underscores_in_headers"] = "Counter"
	MetricTypes["http1__response_flood"] = "Counter"
	MetricTypes["lb_healthy_panic"] = "Counter"
	MetricTypes["lb_local_cluster_not_ok"] = "Counter"
	MetricTypes["lb_recalculate_zone_structures"] = "Counter"
	MetricTypes["lb_subsets_active"] = "Counter"
	MetricTypes["lb_subsets_created"] = "Counter"
	MetricTypes["lb_subsets_fallback"] = "Counter"
	MetricTypes["lb_subsets_fallback_panic"] = "Counter"
	MetricTypes["lb_subsets_removed"] = "Counter"
	MetricTypes["lb_subsets_selected"] = "Counter"
	MetricTypes["lb_zone_cluster_too_small"] = "Counter"
	MetricTypes["lb_zone_no_capacity_left"] = "Counter"
	MetricTypes["lb_zone_number_differs"] = "Counter"
	MetricTypes["lb_zone_routing_all_directly"] = "Counter"
	MetricTypes["lb_zone_routing_cross_zone"] = "Counter"
	MetricTypes["lb_zone_routing_sampled"] = "Counter"
	MetricTypes["max_host_weight"] = "Counter"
	MetricTypes["membership_change"] = "Counter"
	MetricTypes["membership_degraded"] = "Gauge"
	MetricTypes["membership_excluded"] = "Gauge"
	MetricTypes["embership_healthy"] = "Gauge"
	MetricTypes["mMembership_total"] = "Gauge"
	MetricTypes["original_dst_host_invalid"] = "Counter"
	MetricTypes["retry_or_shadow_abandoned"] = "Counter"
	MetricTypes["update_attempt"] = "Counter"
	MetricTypes["update_empty"] = "Counter"
	MetricTypes["update_failure"] = "Counter"
	MetricTypes["update_no_rebuild"] = "Counter"
	MetricTypes["update_success"] = "Counter"
	MetricTypes["upstream_cx_active"] = "Gauge"
	MetricTypes["upstream_cx_close_notify"] = "Counter"
	MetricTypes["upstream_cx_connect_attempts_exceeded"] = "Counter"
	MetricTypes["upstream_cx_connect_fail"] = "Counter"
	MetricTypes["upstream_cx_connect_timeout"] = "Counter"
	MetricTypes["upstream_cx_connect_with_0_rtt"] = "Counter"
	MetricTypes["upstream_cx_destroy"] = "Counter"
	MetricTypes["upstream_cx_destroy_local"] = "Counter"
	MetricTypes["upstream_cx_destroy_local_with_active_rq"] = "Counter"
	MetricTypes["upstream_cx_destroy_remote"] = "Counter"
	MetricTypes["upstream_cx_destroy_remote_with_active_rq"] = "Counter"
	MetricTypes["upstream_cx_destroy_with_active_rq"] = "Counter"
	MetricTypes["upstream_cx_http1_total"] = "Counter"
	MetricTypes["upstream_cx_http2_total"] = "Counter"
	MetricTypes["upstream_cx_http3_total"] = "Counter"
	MetricTypes["upstream_cx_idle_timeout"] = "Counter"
	MetricTypes["upstream_cx_max_duration_reached"] = "Counter"
	MetricTypes["upstream_cx_max_requests"] = "Counter"
	MetricTypes["upstream_cx_none_healthy"] = "Counter"
	MetricTypes["upstream_cx_overflow"] = "Counter"
	MetricTypes["upstream_cx_pool_overflow"] = "Counter"
	MetricTypes["upstream_cx_protocol_error"] = "Counter"
	MetricTypes["upstream_cx_rx_bytes_buffered"] = "Gauge"
	MetricTypes["upstream_cx_rx_bytes_total"] = "Counter"
	MetricTypes["upstream_cx_total"] = "Counter"
	MetricTypes["upstream_cx_tx_bytes_buffered"] = "Gauge"
	MetricTypes["upstream_cx_tx_bytes_total"] = "Counter"
	MetricTypes["upstream_flow_control_backed_up_total"] = "Counter"
	MetricTypes["upstream_flow_control_drained_total"] = "Counter"
	MetricTypes["upstream_flow_control_paused_reading_total"] = "Counter"
	MetricTypes["upstream_flow_control_resumed_reading_total"] = "Counter"
	MetricTypes["upstream_http3_broken"] = "Counter"
	MetricTypes["upstream_internal_redirect_failed_total"] = "Counter"
	MetricTypes["upstream_internal_redirect_succeeded_total"] = "Counter"
	MetricTypes["upstream_rq_0rtt"] = "Counter"
	MetricTypes["upstream_rq_active"] = "Gauge"
	MetricTypes["upstream_rq_cancelled"] = "Counter"
	MetricTypes["upstream_rq_completed"] = "Counter"
	MetricTypes["upstream_rq_maintenance_mode"] = "Counter"
	MetricTypes["upstream_rq_max_duration_reached"] = "Counter"
	MetricTypes["upstream_rq_pending_active"] = "Gauge"
	MetricTypes["upstream_rq_pending_failure_eject"] = "Counter"
	MetricTypes["upstream_rq_pending_overflow"] = "Counter"
	MetricTypes["upstream_rq_pending_total"] = "Counter"
	MetricTypes["upstream_rq_per_try_idle_timeout"] = "Counter"
	MetricTypes["upstream_rq_per_try_timeout"] = "Counter"
	MetricTypes["upstream_rq_retry"] = "Counter"
	MetricTypes["upstream_rq_retry_backoff_exponential"] = "Counter"
	MetricTypes["upstream_rq_retry_backoff_ratelimited"] = "Counter"
	MetricTypes["upstream_rq_retry_limit_exceeded"] = "Counter"
	MetricTypes["upstream_rq_retry_overflow"] = "Counter"
	MetricTypes["upstream_rq_retry_success"] = "Counter"
	MetricTypes["upstream_rq_rx_reset"] = "Counter"
	MetricTypes["upstream_rq_timeout"] = "Counter"
	MetricTypes["upstream_rq_total"] = "Counter"
	MetricTypes["upstream_rq_tx_reset"] = "Counter"
	MetricTypes["version"] = "Gauge"
	return MetricTypes
}
