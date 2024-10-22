import React, { useState, useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { updatePaper, getPaperById } from "../api/api";

const EditPaper: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();

  const [title, setTitle] = useState<string>("");
  const [description, setDescription] = useState<string>("");
  const [url, setUrl] = useState<string>("");
  const [isRead, setIsRead] = useState<boolean>(false);

  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string>("");
  const [success, setSuccess] = useState<string>("");

  useEffect(() => {
    const fetchPaper = async () => {
      if (!id) {
        setError("No paper ID provided.");
        return;
      }

      try {
        const response = await getPaperById(Number(id));
        const paper = response.data.data;
        setTitle(paper.title);
        setDescription(paper.description);
        setUrl(paper.url);
        setIsRead(paper.is_read);
      } catch (err: any) {
        console.error(err);
        setError("Failed to fetch paper details.");
      }
    };

    fetchPaper();
  }, [id]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    setError("");
    setSuccess("");

    if (!title.trim()) {
      setError("Title is required.");
      return;
    }

    if (!url.trim()) {
      setError("URL is required.");
      return;
    }

    const updatedPaper = {
      title,
      description,
      url,
      is_read: isRead,
    };

    try {
      setLoading(true);
      await updatePaper(Number(id), updatedPaper);
      setSuccess("Paper updated successfully!");
      setTimeout(() => {
        navigate("/");
      }, 1500);
    } catch (err: any) {
      console.error(err);
      setError("Failed to update paper. Please try again.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="container mx-auto p-4 max-w-md">
      <h2 className="text-2xl font-bold mb-6 text-gray-800">Edit Paper</h2>

      {success && (
        <div className="bg-green-100 text-green-700 p-3 rounded mb-4">
          {success}
        </div>
      )}

      {error && (
        <div className="bg-red-100 text-red-700 p-3 rounded mb-4">{error}</div>
      )}

      <form
        onSubmit={handleSubmit}
        className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4"
      >
        <div className="mb-4">
          <label
            htmlFor="title"
            className="block text-gray-700 text-sm font-bold mb-2"
          >
            Title<span className="text-red-500">*</span>
          </label>
          <input
            id="title"
            type="text"
            placeholder="Enter paper title"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            required
          />
        </div>

        <div className="mb-4">
          <label
            htmlFor="description"
            className="block text-gray-700 text-sm font-bold mb-2"
          >
            Description
          </label>
          <textarea
            id="description"
            placeholder="Enter paper description"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            rows={4}
          ></textarea>
        </div>

        <div className="mb-4">
          <label
            htmlFor="url"
            className="block text-gray-700 text-sm font-bold mb-2"
          >
            URL<span className="text-red-500">*</span>
          </label>
          <input
            id="url"
            type="url"
            placeholder="https://example.com"
            value={url}
            onChange={(e) => setUrl(e.target.value)}
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            required
          />
        </div>

        <div className="mb-6 flex items-center">
          <input
            id="is_read"
            type="checkbox"
            checked={isRead}
            onChange={(e) => setIsRead(e.target.checked)}
            className="mr-2 leading-tight"
          />
          <label htmlFor="is_read" className="text-gray-700 text-sm font-bold">
            Mark as Read
          </label>
        </div>

        <div className="flex items-center justify-start">
          <button
            type="submit"
            disabled={loading}
            className={`${
              loading
                ? "bg-gray-400 cursor-not-allowed"
                : "bg-gray-800 hover:bg-gray-700"
            } text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline`}
          >
            {loading ? "Updating..." : "Update Paper"}
          </button>
        </div>
      </form>
    </div>
  );
};

export default EditPaper;
