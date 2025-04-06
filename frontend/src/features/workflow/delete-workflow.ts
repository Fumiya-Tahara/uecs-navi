// import { deleteWorkflowWithUI } from "../api/mocks/workflow-api";
import { deleteWorkflowWithUI } from "../api/workflow/delete-workflow";

export async function deleteWorkflow(workflowID: number) {
  deleteWorkflowWithUI(workflowID);

  try {
    await deleteWorkflowWithUI(workflowID);
  } catch (error) {
    console.error("Failed to delete workflow:", error);
  }
}
