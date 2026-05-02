// test/__tests__/address.test.js
// Unit API Test untuk Address Endpoints

const axios = require("axios");
const { BASE_URL } = require("../config");

const api = axios.create({
  baseURL: BASE_URL,
  validateStatus: () => true,
});

let authToken = "";
const timestamp = Date.now();
const testUser = {
  name: `Address Tester ${timestamp}`,
  email: `addresstest${timestamp}@example.com`,
  phone: "082200000002",
  password: "testpass123",
};

beforeAll(async () => {
  await api.post("/api/auth/register", testUser);
  const loginRes = await api.post("/api/auth/login", {
    email: testUser.email,
    password: testUser.password,
  });
  authToken = loginRes.data?.data?.token || "";
});

describe("Address API", () => {
  // ===========================================================================
  // GET MY ADDRESSES - Protected
  // ===========================================================================
  describe("GET /api/address/", () => {
    it("harus mengembalikan 401 tanpa token", async () => {
      const res = await api.get("/api/address/");
      expect([401, 403]).toContain(res.status);
    });

    it("harus mengembalikan daftar alamat milik user", async () => {
      const res = await api.get("/api/address/", {
        headers: { Authorization: `Bearer ${authToken}` },
      });

      expect(res.status).toBe(200);
      expect(res.data.success).toBe(true);
      expect(Array.isArray(res.data.data)).toBe(true);
    });
  });

  // ===========================================================================
  // CREATE ADDRESS - Protected
  // ===========================================================================
  describe("POST /api/address/", () => {
    it("harus mengembalikan 401 tanpa token", async () => {
      const res = await api.post("/api/address/", {
        title: "Rumah",
        address_details: "Jl. Test No. 1",
      });
      expect([401, 403]).toContain(res.status);
    });

    it("harus berhasil membuat alamat baru", async () => {
      const res = await api.post(
        "/api/address/",
        {
          title: "Rumah",
          address_details: "Jl. Merdeka No. 17, Jakarta Pusat",
          prov_id: "31",
          city_id: "151",
          is_primary: true,
        },
        { headers: { Authorization: `Bearer ${authToken}` } }
      );

      expect(res.status).toBe(201);
      expect(res.data.success).toBe(true);
      expect(res.data.data).toHaveProperty("id");
    });

    it("harus bisa membuat lebih dari satu alamat", async () => {
      const res = await api.post(
        "/api/address/",
        {
          title: "Kantor",
          address_details: "Jl. Sudirman Kav. 55, Jakarta Selatan",
          prov_id: "31",
          city_id: "151",
          is_primary: false,
        },
        { headers: { Authorization: `Bearer ${authToken}` } }
      );

      expect(res.status).toBe(201);
      expect(res.data.success).toBe(true);
    });
  });
});
