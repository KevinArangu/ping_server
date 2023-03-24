export interface StatsResponse {
    total_pings: string;
    total_local_pings: string;
    completed_local_pings: string;
    error_local_pings: string;
    total_remote_pings: string;
    completed_remote_pings: string;
    error_remote_pings: string;
    is_local_conected: boolean;
    is_remote_conected: boolean;
}