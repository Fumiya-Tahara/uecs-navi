import {
  createContext,
  useContext,
  useState,
  Dispatch,
  SetStateAction,
} from "react";

type M304IDsState = {
  m304IDs: number[] | null;
};

type m304IDsContextType = [
  M304IDsState,
  Dispatch<SetStateAction<M304IDsState>>
];

const M304IDsContext = createContext<m304IDsContextType>([
  {
    m304IDs: null,
  },
  () => {},
]);

export const M304IDsProvider = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  const [m304IDs, setM304IDs] = useState<M304IDsState>({
    m304IDs: [],
  });

  return (
    <M304IDsContext.Provider value={[m304IDs, setM304IDs]}>
      {children}
    </M304IDsContext.Provider>
  );
};

export const useM304IDs = () => {
  return useContext(M304IDsContext);
};

export default M304IDsContext;
