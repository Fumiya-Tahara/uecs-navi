import { WorkflowWithUIResponse } from "@/types/api";
import {
  createContext,
  useContext,
  useState,
  Dispatch,
  SetStateAction,
} from "react";

type selectedDataState = {
  selectedM304ID: number | null;
  selectedWorkflow: WorkflowWithUIResponse | null;
};

type selectedDataContextType = [
  selectedDataState,
  Dispatch<SetStateAction<selectedDataState>>
];

const SelectedDataContext = createContext<selectedDataContextType>([
  {
    selectedM304ID: null,
    selectedWorkflow: null,
  },
  () => {},
]);

export const SelectedDataProvider = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  const [selectedData, setSelectedData] = useState<selectedDataState>({
    selectedM304ID: null,
    selectedWorkflow: null,
  });

  return (
    <SelectedDataContext.Provider value={[selectedData, setSelectedData]}>
      {children}
    </SelectedDataContext.Provider>
  );
};

export const useSelectedData = () => {
  return useContext(SelectedDataContext);
};

export default SelectedDataContext;
