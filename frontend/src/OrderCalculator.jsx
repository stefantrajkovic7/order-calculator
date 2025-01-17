import { useState } from "react";
import { calculatePacks } from "./api";

function OrderCalculator() {
  const [order, setOrder] = useState("");
  const [result, setResult] = useState(null);
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError("");
    setResult(null);

    try {
      const response = await calculatePacks(Number(order));
      setResult(response.packs);
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="p-6 max-w-lg mx-auto bg-white rounded-lg shadow-md">
      <h1 className="text-2xl font-bold text-gray-800 mb-4 text-center">Calculate Packs for Order</h1>
      <form onSubmit={handleSubmit} className="space-y-4">
        <label className="block">
          <span className="text-gray-700">Items:</span>
          <input
            type="number"
            className="w-full mt-1 p-2 border rounded focus:ring-2 focus:ring-blue-500"
            value={order}
            onChange={(e) => setOrder(e.target.value)}
            placeholder="Enter number of items"
          />
        </label>
        <button
          type="submit"
          className="w-full p-2 bg-blue-500 text-white rounded hover:bg-blue-600 transition"
          disabled={loading}
        >
          {loading ? "Calculating..." : "Calculate"}
        </button>
      </form>
      {error && <p className="text-red-500 mt-4">{error}</p>}
      {result && (
        <div className="mt-4">
          <h2 className="text-lg font-bold text-gray-700">Results:</h2>
          <table className="table-auto w-full mt-2 border border-collapse border-gray-300">
            <thead>
              <tr className="bg-gray-100">
                <th className="border px-4 py-2">Pack</th>
                <th className="border px-4 py-2">Quantity</th>
              </tr>
            </thead>
            <tbody>
              {Object.entries(result).map(([pack, quantity]) => (
                <tr key={pack}>
                  <td className="border px-4 py-2 text-center">{pack}</td>
                  <td className="border px-4 py-2 text-center">{quantity}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
}

export default OrderCalculator;