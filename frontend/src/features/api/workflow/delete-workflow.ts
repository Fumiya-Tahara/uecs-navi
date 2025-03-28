import { apiClient } from "@/lib/api-client";

// 戻り値は作成したレコードのID
export async function deleteWorkflowWithUI(
  workflowID: number
): Promise<number> {
  try {
    const response = await apiClient.delete(`/workflows-with-ui/${workflowID}`);
    return response.data;
  } catch (error) {
    console.error("Error deleting workflow:", error);
    throw error;
  }
}
