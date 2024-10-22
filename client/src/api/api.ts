import axios from "axios";

const API_URL = import.meta.env.VITE_API_URL;

const api = axios.create({
  baseURL: API_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

export interface Paper {
  id: number;
  title: string;
  description: string;
  url: string;
  is_read: boolean;
  created_at: string;
  updated_at: string;
  deleted_at: string;
}

export const getPapers = () => api.get("/papers");
export const getPaperById = (id: number) => api.get(`/papers/${id}`);
export const createPaper = (data: Paper) => api.post("/papers", data);
export const updatePaper = (id: number, data: Paper) =>
  api.patch(`/papers/${id}`, data);
export const deletePaper = (id: number) => api.delete(`/papers/${id}`);
