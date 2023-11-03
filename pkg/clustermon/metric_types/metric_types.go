package metric_types

// NewMetricTypes parse the metric types from a yaml file
func NewMetricTypes() map[string]string {
	MetricTypes := make(map[string]string)
	MetricTypes["rto"] = "Gauge"
	MetricTypes["rtt"] = "Gauge"
	MetricTypes["delivery_rate"] = "Gauge"
	MetricTypes["bytes_send"] = "Gauge"
	MetricTypes["bytes_received"] = "Gauge"
	MetricTypes["min_rtt"] = "Gauge"
	MetricTypes["assignment_stale"] = "Gauge"
	MetricTypes["assignment_timeout_received"] = "Gauge"
	MetricTypes["bind_errors"] = "Gauge"
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
	MetricTypes["default__total_match_count"] = "Gauge"
	MetricTypes["health_check__attempt"] = "Gauge"
	MetricTypes["health_check__degraded"] = "Gauge"
	MetricTypes["health_check__failure"] = "Gauge"
	MetricTypes["health_check__healthy"] = "Gauge"
	MetricTypes["health_check__network_failure"] = "Gauge"
	MetricTypes["health_check__passive_failure"] = "Gauge"
	MetricTypes["health_check__success"] = "Gauge"
	MetricTypes["health_check__verify_cluster"] = "Gauge"
	MetricTypes["http1__dropped_headers_with_underscores"] = "Gauge"
	MetricTypes["http1__metadata_not_supported_error"] = "Gauge"
	MetricTypes["http1__requests_rejected_with_underscores_in_headers"] = "Gauge"
	MetricTypes["http1__response_flood"] = "Gauge"
	MetricTypes["lb_healthy_panic"] = "Gauge"
	MetricTypes["lb_local_cluster_not_ok"] = "Gauge"
	MetricTypes["lb_recalculate_zone_structures"] = "Gauge"
	MetricTypes["lb_subsets_active"] = "Gauge"
	MetricTypes["lb_subsets_created"] = "Gauge"
	MetricTypes["lb_subsets_fallback"] = "Gauge"
	MetricTypes["lb_subsets_fallback_panic"] = "Gauge"
	MetricTypes["lb_subsets_removed"] = "Gauge"
	MetricTypes["lb_subsets_selected"] = "Gauge"
	MetricTypes["lb_zone_cluster_too_small"] = "Gauge"
	MetricTypes["lb_zone_no_capacity_left"] = "Gauge"
	MetricTypes["lb_zone_number_differs"] = "Gauge"
	MetricTypes["lb_zone_routing_all_directly"] = "Gauge"
	MetricTypes["lb_zone_routing_cross_zone"] = "Gauge"
	MetricTypes["lb_zone_routing_sampled"] = "Gauge"
	MetricTypes["max_host_weight"] = "Gauge"
	MetricTypes["membership_change"] = "Gauge"
	MetricTypes["membership_degraded"] = "Gauge"
	MetricTypes["membership_excluded"] = "Gauge"
	MetricTypes["embership_healthy"] = "Gauge"
	MetricTypes["mMembership_total"] = "Gauge"
	MetricTypes["original_dst_host_invalid"] = "Gauge"
	MetricTypes["retry_or_shadow_abandoned"] = "Gauge"
	MetricTypes["update_attempt"] = "Gauge"
	MetricTypes["update_empty"] = "Gauge"
	MetricTypes["update_failure"] = "Gauge"
	MetricTypes["update_no_rebuild"] = "Gauge"
	MetricTypes["update_success"] = "Gauge"
	MetricTypes["upstream_cx_active"] = "Gauge"
	MetricTypes["upstream_cx_close_notify"] = "Gauge"
	MetricTypes["upstream_cx_connect_attempts_exceeded"] = "Gauge"
	MetricTypes["upstream_cx_connect_fail"] = "Gauge"
	MetricTypes["upstream_cx_connect_timeout"] = "Gauge"
	MetricTypes["upstream_cx_connect_with_0_rtt"] = "Gauge"
	MetricTypes["upstream_cx_destroy"] = "Gauge"
	MetricTypes["upstream_cx_destroy_local"] = "Gauge"
	MetricTypes["upstream_cx_destroy_local_with_active_rq"] = "Gauge"
	MetricTypes["upstream_cx_destroy_remote"] = "Gauge"
	MetricTypes["upstream_cx_destroy_remote_with_active_rq"] = "Gauge"
	MetricTypes["upstream_cx_destroy_with_active_rq"] = "Gauge"
	MetricTypes["upstream_cx_http1_total"] = "Gauge"
	MetricTypes["upstream_cx_http2_total"] = "Gauge"
	MetricTypes["upstream_cx_http3_total"] = "Gauge"
	MetricTypes["upstream_cx_idle_timeout"] = "Gauge"
	MetricTypes["upstream_cx_max_duration_reached"] = "Gauge"
	MetricTypes["upstream_cx_max_requests"] = "Gauge"
	MetricTypes["upstream_cx_none_healthy"] = "Gauge"
	MetricTypes["upstream_cx_overflow"] = "Gauge"
	MetricTypes["upstream_cx_pool_overflow"] = "Gauge"
	MetricTypes["upstream_cx_protocol_error"] = "Gauge"
	MetricTypes["upstream_cx_rx_bytes_buffered"] = "Gauge"
	MetricTypes["upstream_cx_rx_bytes_total"] = "Gauge"
	MetricTypes["upstream_cx_total"] = "Gauge"
	MetricTypes["upstream_cx_tx_bytes_buffered"] = "Gauge"
	MetricTypes["upstream_cx_tx_bytes_total"] = "Gauge"
	MetricTypes["upstream_flow_control_backed_up_total"] = "Gauge"
	MetricTypes["upstream_flow_control_drained_total"] = "Gauge"
	MetricTypes["upstream_flow_control_paused_reading_total"] = "Gauge"
	MetricTypes["upstream_flow_control_resumed_reading_total"] = "Gauge"
	MetricTypes["upstream_http3_broken"] = "Gauge"
	MetricTypes["upstream_internal_redirect_failed_total"] = "Gauge"
	MetricTypes["upstream_internal_redirect_succeeded_total"] = "Gauge"
	MetricTypes["upstream_rq_0rtt"] = "Gauge"
	MetricTypes["upstream_rq_active"] = "Gauge"
	MetricTypes["upstream_rq_cancelled"] = "Gauge"
	MetricTypes["upstream_rq_completed"] = "Gauge"
	MetricTypes["upstream_rq_maintenance_mode"] = "Gauge"
	MetricTypes["upstream_rq_max_duration_reached"] = "Gauge"
	MetricTypes["upstream_rq_pending_active"] = "Gauge"
	MetricTypes["upstream_rq_pending_failure_eject"] = "Gauge"
	MetricTypes["upstream_rq_pending_overflow"] = "Gauge"
	MetricTypes["upstream_rq_pending_total"] = "Gauge"
	MetricTypes["upstream_rq_per_try_idle_timeout"] = "Gauge"
	MetricTypes["upstream_rq_per_try_timeout"] = "Gauge"
	MetricTypes["upstream_rq_retry"] = "Gauge"
	MetricTypes["upstream_rq_retry_backoff_exponential"] = "Gauge"
	MetricTypes["upstream_rq_retry_backoff_ratelimited"] = "Gauge"
	MetricTypes["upstream_rq_retry_limit_exceeded"] = "Gauge"
	MetricTypes["upstream_rq_retry_overflow"] = "Gauge"
	MetricTypes["upstream_rq_retry_success"] = "Gauge"
	MetricTypes["upstream_rq_rx_reset"] = "Gauge"
	MetricTypes["upstream_rq_timeout"] = "Gauge"
	MetricTypes["upstream_rq_total"] = "Gauge"
	MetricTypes["upstream_rq_tx_reset"] = "Gauge"
	MetricTypes["version"] = "Gauge"
	MetricTypes["retrain_rate"] = "Gauge"
	return MetricTypes
}