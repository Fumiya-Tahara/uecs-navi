import { apiClient } from "@/lib/api-client";

export async function getM304s(): Promise<number[]> {
  try {
    const response = await apiClient.get(`/m304s`);
    return response.data;
  } catch (error) {
    console.error("Error getting M304s:", error);
    throw error;
  }
}
