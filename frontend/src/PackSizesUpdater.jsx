import { useState } from "react";
import { updatePackSizes } from "./api";

function PackSizesUpdater() {
  const [sizes, setSizes] = useState("");
  const [message, setMessage] = useState("");
  const [error, setError] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const sizeArray = sizes.split(",").map(Number);
      await updatePackSizes(sizeArray);
      setMessage("Pack sizes updated successfully!");
      setError("");
    } catch (err) {
      setError(err.message);
    }
  };

  return (
    <div>
      <h1>Update Pack Sizes</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          value={sizes}
          onChange={(e) => setSizes(e.target.value)}
          placeholder="Enter pack sizes (comma-separated)"
        />
        <button type="submit">Update</button>
      </form>
      {message && <p>{message}</p>}
      {error && <p>{error}</p>}
    </div>
  );
}

export default PackSizesUpdater;