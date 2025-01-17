import { useState } from "react";
import { calculatePacks } from "./api";

function OrderCalculator() {
  const [order, setOrder] = useState("");
  const [result, setResult] = useState(null);
  const [error, setError] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const data = await calculatePacks(Number(order));
      setResult(data.packs);
      setError("");
    } catch (err) {
      setError(err.message);
    }
  };

  return (
    <div>
      <h1>Order Calculator</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="number"
          value={order}
          onChange={(e) => setOrder(e.target.value)}
          placeholder="Enter order quantity"
        />
        <button type="submit">Calculate</button>
      </form>
      {error && <p>{error}</p>}
      {result && (
        <ul>
          {Object.entries(result).map(([pack, qty]) => (
            <li key={pack}>{`${qty} x ${pack}`}</li>
          ))}
        </ul>
      )}
    </div>
  );
}

export default OrderCalculator;