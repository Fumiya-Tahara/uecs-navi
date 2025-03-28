import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import App from "@/app/App.tsx";
import { M304IDsProvider } from "./hooks/m304ids-context";
// import './index.css'

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <M304IDsProvider>
      <App />
    </M304IDsProvider>
  </StrictMode>
);
