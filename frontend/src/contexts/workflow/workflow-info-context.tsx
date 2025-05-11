import { DeviceResponse, WorkflowWithUIResponse } from "@/types/api";
import {
  createContext,
  useContext,
  useState,
  Dispatch,
  SetStateAction,
} from "react";

type workflowInfoState = {
  m304DeviceMap: Map<number, DeviceResponse[]>;
  m304WorkflowMap: Map<number, WorkflowWithUIResponse[]>;
};

type workflowInfoContextType = [
  workflowInfoState,
  Dispatch<SetStateAction<workflowInfoState>>
];

const WorkflowInfoContext = createContext<workflowInfoContextType>([
  {
    m304DeviceMap: new Map<number, DeviceResponse[]>(),
    m304WorkflowMap: new Map<number, WorkflowWithUIResponse[]>(),
  },
  () => {},
]);

export const WorkflowInfoProvider = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  const [workflowInfo, setWorkflowInfo] = useState<workflowInfoState>({
    m304DeviceMap: new Map<number, DeviceResponse[]>(),
    m304WorkflowMap: new Map<number, WorkflowWithUIResponse[]>(),
  });

  return (
    <WorkflowInfoContext.Provider value={[workflowInfo, setWorkflowInfo]}>
      {children}
    </WorkflowInfoContext.Provider>
  );
};

export const useWorkflowInfo = () => {
  return useContext(WorkflowInfoContext);
};

export default WorkflowInfoContext;
