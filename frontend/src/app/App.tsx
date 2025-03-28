import { useM304IDs } from "@/hooks/m304ids-context";
import { AppRouter } from "./routes";
import { useEffect } from "react";
import { getM304s } from "@/features/api/mocks/utility-api";

function App() {
  const [m304IDs, setM304IDs] = useM304IDs();

  useEffect(() => {
    const m304IDsList: number[] = getM304s();

    m304IDs.m304IDs = m304IDsList;
    setM304IDs(m304IDs);
  }, []);
  return (
    <>
      <AppRouter />
    </>
  );
}

export default App;
