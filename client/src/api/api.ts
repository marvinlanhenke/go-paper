import axios from "axios";

const API_URL = import.meta.env.VITE_API_URL;

const api = axios.create({
  baseURL: API_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

export interface CreatePaper {
  title: string;
  description: string;
  url: string;
}

export interface UpdatePaper extends CreatePaper {
  is_read: boolean;
}

export interface Paper extends UpdatePaper {
  id: number;
}

export const getPapers = () => api.get("/papers");
export const getPaperById = (id: number) => api.get(`/papers/${id}`);
export const createPaper = (data: CreatePaper) => api.post("/papers", data);
export const updatePaper = (id: number, data: UpdatePaper) =>
  api.patch(`/papers/${id}`, data);
export const deletePaper = (id: number) => api.delete(`/papers/${id}`);
