import { Workflow } from "@/types/api";
import { apiClient } from "@/lib/api-client";

export async function getWorkflows(): Promise<Workflow[]> {
  try {
    const response = await apiClient.get(`/workflows`);
    return response.data;
  } catch (error) {
    console.error("Error getting workflows:", error);
    throw error;
  }
}
