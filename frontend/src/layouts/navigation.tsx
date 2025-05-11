import { Box, Toolbar, CssBaseline } from "@mui/material";
import { ThemeProvider } from "@mui/material/styles";
import Header from "@/components/header";
import { uecsTheme } from "@/styles/theme";

export function Navigation({ children }: { children: React.ReactNode }) {
  return (
    <ThemeProvider theme={uecsTheme}>
      <Box sx={{ display: "flex" }}>
        <CssBaseline />
        <Header />
        <Box component="main" sx={{ flexGrow: 1, padding: 8 }}>
          <Toolbar />
          {children}
        </Box>
      </Box>
    </ThemeProvider>
  );
}
