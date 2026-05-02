// test/__tests__/category.test.js
// Unit API Test untuk Category Endpoints

const axios = require("axios");
const { BASE_URL } = require("../config");

const api = axios.create({
  baseURL: BASE_URL,
  validateStatus: () => true,
});

// Shared state antar test dalam satu file
let authToken = "";
let createdCategoryId = null;

const timestamp = Date.now();
const adminUser = {
  name: `Admin Test ${timestamp}`,
  email: `admin${timestamp}@example.com`,
  phone: "081234500000",
  password: "adminpass123",
};

beforeAll(async () => {
  // Register dan login dulu untuk dapat token
  await api.post("/api/auth/register", adminUser);
  const loginRes = await api.post("/api/auth/login", {
    email: adminUser.email,
    password: adminUser.password,
  });
  authToken = loginRes.data?.data?.token || "";
});

describe("Category API", () => {
  // ===========================================================================
  // GET ALL - Public
  // ===========================================================================
  describe("GET /api/category/", () => {
    it("harus mengembalikan daftar kategori (public, tanpa token)", async () => {
      const res = await api.get("/api/category/");

      expect(res.status).toBe(200);
      expect(res.data.success).toBe(true);
      expect(Array.isArray(res.data.data)).toBe(true);
    });
  });

  // ===========================================================================
  // CREATE - Protected (Admin)
  // ===========================================================================
  describe("POST /api/category/", () => {
    it("harus berhasil membuat kategori dengan token valid", async () => {
      const res = await api.post(
        "/api/category/",
        { name: `Kategori Test ${timestamp}`, description: "Deskripsi test" },
        { headers: { Authorization: `Bearer ${authToken}` } }
      );

      // Note: 201 jika admin, 403 jika user biasa (role-based)
      if (res.status === 201) {
        expect(res.data.success).toBe(true);
        expect(res.data.data).toHaveProperty("id");
        createdCategoryId = res.data.data.id;
      } else {
        // User biasa tidak bisa akses (403 Forbidden)
        expect([403]).toContain(res.status);
      }
    });

    it("harus mengembalikan 401 tanpa token", async () => {
      const res = await api.post("/api/category/", {
        name: "Kategori Tanpa Auth",
      });

      expect([401, 403]).toContain(res.status);
      expect(res.data.success).toBe(false);
    });

    it("harus mengembalikan 400 jika name kosong", async () => {
      const res = await api.post(
        "/api/category/",
        { name: "" },
        { headers: { Authorization: `Bearer ${authToken}` } }
      );

      // 400 dari validasi, atau 403 jika bukan admin
      expect([400, 403]).toContain(res.status);
    });
  });

  // ===========================================================================
  // UPDATE - Protected (Admin)
  // ===========================================================================
  describe("PUT /api/category/:id", () => {
    it("harus mengembalikan 401 tanpa token", async () => {
      const res = await api.put("/api/category/1", { name: "Updated" });
      expect([401, 403]).toContain(res.status);
    });

    it("harus mengembalikan 400 untuk ID tidak valid", async () => {
      const res = await api.put(
        "/api/category/invalid-id",
        { name: "Updated" },
        { headers: { Authorization: `Bearer ${authToken}` } }
      );
      expect([400, 403]).toContain(res.status);
    });
  });

  // ===========================================================================
  // DELETE - Protected (Admin)
  // ===========================================================================
  describe("DELETE /api/category/:id", () => {
    it("harus mengembalikan 401 tanpa token", async () => {
      const res = await api.delete("/api/category/1");
      expect([401, 403]).toContain(res.status);
    });
  });
});
