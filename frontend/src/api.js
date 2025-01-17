// simplifying the api flow, this type of const would be better in a .env file
const API_BASE_URL = "http://localhost:8080";

export const calculatePacks = async (order) => {
  const response = await fetch(`${API_BASE_URL}/calculate`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ order }),
  });
  if (!response.ok) throw new Error("Error calculating packs");
  return response.json();
};

export const updatePackSizes = async (sizes) => {
  const response = await fetch(`${API_BASE_URL}/packs`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ pack_sizes: sizes }),
  });
  if (!response.ok) throw new Error("Error updating pack sizes");
  return response.json();
};