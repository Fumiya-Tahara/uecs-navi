import { WorkflowWithUIResponse } from "@/types/api";
import { apiClient } from "@/lib/api-client";

export async function getWorkflows(
  m304ID: number
): Promise<WorkflowWithUIResponse[]> {
  try {
    const response = await apiClient.get(`/workflows-with-ui/${m304ID}`);
    return response.data;
  } catch (error) {
    console.error("Error getting workflows:", error);
    throw error;
  }
}
