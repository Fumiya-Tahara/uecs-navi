import { deleteWorkflowWithUI } from "../api/mocks/workflow-api";

export function deleteWorkflow(workflowID: number) {
  deleteWorkflowWithUI(workflowID);
}
