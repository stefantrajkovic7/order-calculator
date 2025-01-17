import { useState } from "react";
import { updatePackSizes } from "./api";

function PackSizesUpdater() {
  const [sizes, setSizes] = useState("");
  const [message, setMessage] = useState("");
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setMessage("");
    setError("");

    try {
      const sizesArray = sizes.split(",").map(Number);
      await updatePackSizes(sizesArray);
      setMessage("Pack sizes updated successfully!");
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="p-6 max-w-lg mx-auto bg-white rounded-lg shadow-md mt-8">
      <h1 className="text-2xl font-bold text-gray-800 mb-4 text-center">Update Pack Sizes</h1>
      <form onSubmit={handleSubmit} className="space-y-4">
        <label className="block">
          <span className="text-gray-700">Pack Sizes:</span>
          <input
            type="text"
            className="w-full mt-1 p-2 border rounded focus:ring-2 focus:ring-green-500"
            value={sizes}
            onChange={(e) => setSizes(e.target.value)}
            placeholder="Enter sizes (comma-separated)"
          />
        </label>
        <button
          type="submit"
          className="w-full p-2 bg-green-500 text-white rounded hover:bg-green-600 transition"
          disabled={loading}
        >
          {loading ? "Updating..." : "Update"}
        </button>
      </form>
      {error && <p className="text-red-500 mt-4">{error}</p>}
      {message && <p className="text-green-500 mt-4">{message}</p>}
    </div>
  );
}

export default PackSizesUpdater;