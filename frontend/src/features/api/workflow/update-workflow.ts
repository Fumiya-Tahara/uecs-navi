import { WorkflowWithUIRequest } from "@/types/api";
import { apiClient } from "@/lib/api-client";

// 戻り値は更新したレコードのID
export async function updateWorkflowWithUI(
  data: WorkflowWithUIRequest
): Promise<number> {
  try {
    const response = await apiClient.put(`/workflows-with-ui`, data);
    return response.data;
  } catch (error) {
    console.error("Error updating workflow:", error);
    throw error;
  }
}
