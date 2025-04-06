import { useM304IDs } from "@/hooks/m304ids-context";
import { AppRouter } from "./routes";
import { useEffect } from "react";
// import { getM304s } from "@/features/api/mocks/utility-api";
import { getM304s } from "@/features/api/m304s/get-m304s";

function App() {
  const [, setM304IDs] = useM304IDs();

  useEffect(() => {
    const fetchM304IDs = async () => {
      try {
        const m304IDsList: number[] = await getM304s();
        setM304IDs({ m304IDs: m304IDsList });
      } catch (error) {
        console.error("Failed to fetch M304 IDs:", error);
      }
    };

    fetchM304IDs();
  }, []);
  return (
    <>
      <AppRouter />
    </>
  );
}

export default App;
