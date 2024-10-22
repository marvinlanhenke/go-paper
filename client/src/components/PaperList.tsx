import React, { useEffect, useState } from "react";
import { getPapers, deletePaper, Paper } from "../api/api";
import { Link } from "react-router-dom";
import {
  BookOpenIcon,
  DocumentCheckIcon,
  PlusCircleIcon,
} from "@heroicons/react/16/solid";

const PaperList: React.FC = () => {
  const [papers, setPapers] = useState<Paper[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string>("");

  useEffect(() => {
    fetchItems();
  }, []);

  const fetchItems = async () => {
    try {
      const response = await getPapers();
      setPapers(response.data.data);
    } catch (err) {
      setError("failed to fetch papers");
    } finally {
      setLoading(false);
    }
  };

  const handleDelete = async (id: number) => {
    if (!window.confirm("Are you sure you want to delete this item?")) return;

    try {
      await deletePaper(id);
      setPapers(papers.filter((item) => item.id !== id));
    } catch (err) {
      setError("failed to delete paper");
    }
  };

  if (loading) return <p className="text-center">Loading...</p>;
  if (error) return <p className="text-center text-red-500">{error}</p>;

  return (
    <div className="container mx-auto p-4">
      <div className="flex items-center mb-6">
        <h2 className="text-2xl font-bold text-gray-800">Reading List</h2>
        <Link to="/create">
          <PlusCircleIcon className="w-10 h-10 p-2 hover:text-gray-700" />
        </Link>
      </div>

      <ul className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
        {papers.map((item) => (
          <li
            key={item.id}
            className="bg-white shadow-md rounded p-4 flex flex-col justify-between h-full"
          >
            <div>
              <span className="block text-xl font-semibold mb-2">
                {item.title}
              </span>
              <p className="text-gray-600 mb-2">{item.description}</p>
              <a
                href={item.url}
                target="_blank"
                rel="noopener noreferrer"
                className="text-teal-500 hover:underline"
              >
                {item.url}
              </a>
            </div>

            <div className="flex justify-between items-center mt-4">
              <div className="flex items-center">
                <Link
                  to={`/update/${item.id}`}
                  className="hover:underline mr-4"
                >
                  Edit
                </Link>
                <button
                  onClick={() => handleDelete(item.id)}
                  className="hover:underline"
                >
                  Delete
                </button>
              </div>

              <div>
                {item.is_read ? (
                  <DocumentCheckIcon
                    className="h-6 w-6 text-teal-400"
                    aria-label="Read"
                  />
                ) : (
                  <BookOpenIcon
                    className="h-6 w-6 text-cyan-400"
                    aria-label="Not Read"
                  />
                )}
              </div>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default PaperList;
