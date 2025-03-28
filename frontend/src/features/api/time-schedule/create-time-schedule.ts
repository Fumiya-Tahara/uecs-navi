import { TimeScheduleRequest } from "@/types/api";
import { apiClient } from "@/lib/api-client";

// 戻り値は作成したレコードのID
export async function createTimeSchedule(
  data: TimeScheduleRequest
): Promise<number> {
  try {
    const response = await apiClient.post(`/time-schedule`, data);
    return response.data;
  } catch (error) {
    console.error("Error creating time schedule:", error);
    throw error;
  }
}
