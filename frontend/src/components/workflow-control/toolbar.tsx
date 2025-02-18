import {
  AppBar,
  Box,
  Button,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  Toolbar,
} from "@mui/material";
import Grid from "@mui/material/Grid2";
import { grey } from "@mui/material/colors";

export function SelectToolbar() {
  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar
        position="static"
        sx={{
          backgroundColor: "#FFFFFF",
          color: grey[800],
        }}
      >
        <Toolbar>
          <Grid container spacing={2} sx={{ width: "100%" }}>
            <Grid size={8} sx={{ display: "flex", gap: 1 }}>
              <SelectM304 />
              <SelectWorkflow />
            </Grid>
            <Grid size={4}>
              <SaveWorkflowButton />
            </Grid>
          </Grid>
        </Toolbar>
      </AppBar>
    </Box>
  );
}

function SaveWorkflowButton() {
  return (
    <Button variant="contained" size="small">
      保存する
    </Button>
  );
}

function SelectM304() {
  return (
    <FormControl fullWidth size="small">
      <InputLabel id="demo-simple-select-label">M304</InputLabel>
      <Select
        labelId="demo-simple-select-label"
        id="demo-simple-select"
        label="Age"
      >
        <MenuItem value={10}>Ten</MenuItem>
        <MenuItem value={20}>Twenty</MenuItem>
        <MenuItem value={30}>Thirty</MenuItem>
      </Select>
    </FormControl>
  );
}

function SelectWorkflow() {
  return (
    <FormControl fullWidth size="small">
      <InputLabel id="demo-simple-select-label">ワークフロー</InputLabel>
      <Select
        labelId="demo-simple-select-label"
        id="demo-simple-select"
        label="Age"
      >
        <MenuItem value={10}>Ten</MenuItem>
        <MenuItem value={20}>Twenty</MenuItem>
        <MenuItem value={30}>Thirty</MenuItem>
      </Select>
    </FormControl>
  );
}
