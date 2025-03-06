import { create } from "zustand";
import { Node, Edge } from "@xyflow/react";

type NodesAndEdgesStore = {
  nodes: Node[];
  edges: Edge[];
  setNodes: (nodes: Node[]) => void;
  setEdges: (edges: Edge[]) => void;
};

export const useNodesAndEdges = create<NodesAndEdgesStore>((set) => ({
  nodes: [],
  edges: [],
  setNodes: (nodes) => set({ nodes: nodes }),
  setEdges: (edges) => set({ edges: edges }),
}));
