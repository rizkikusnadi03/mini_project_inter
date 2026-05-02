// test/__tests__/product.test.js
// Unit API Test untuk Product Endpoints

const axios = require("axios");
const FormData = require("form-data");
const { BASE_URL } = require("../config");

const api = axios.create({
  baseURL: BASE_URL,
  validateStatus: () => true,
});

let authToken = "";
const timestamp = Date.now();
const testUser = {
  name: `Product Tester ${timestamp}`,
  email: `producttest${timestamp}@example.com`,
  phone: "082200000001",
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

describe("Product API", () => {
  // ===========================================================================
  // GET ALL - Public
  // ===========================================================================
  describe("GET /api/product/", () => {
    it("harus mengembalikan daftar produk dengan pagination", async () => {
      const res = await api.get("/api/product/");

      expect(res.status).toBe(200);
      expect(res.data.success).toBe(true);
      expect(res.data.data).toHaveProperty("products");
      expect(res.data.data).toHaveProperty("total");
      expect(res.data.data).toHaveProperty("page");
      expect(res.data.data).toHaveProperty("limit");
      expect(Array.isArray(res.data.data.products)).toBe(true);
    });

    it("harus mendukung query parameter page dan limit", async () => {
      const res = await api.get("/api/product/?page=1&limit=5");

      expect(res.status).toBe(200);
      expect(res.data.data.page).toBe(1);
      expect(res.data.data.limit).toBe(5);
    });

    it("harus mendukung query parameter search", async () => {
      const res = await api.get("/api/product/?search=laptop");

      expect(res.status).toBe(200);
      expect(res.data.success).toBe(true);
    });
  });

  // ===========================================================================
  // CREATE - Protected
  // ===========================================================================
  describe("POST /api/product/", () => {
    it("harus mengembalikan 401 tanpa token", async () => {
      const form = new FormData();
      form.append("name", "Test Product");
      form.append("category_id", "1");
      form.append("price", "50000");
      form.append("stock", "10");

      const res = await api.post("/api/product/", form, {
        headers: form.getHeaders(),
      });

      expect([401, 403]).toContain(res.status);
      expect(res.data.success).toBe(false);
    });

    it("harus mengembalikan 400 jika field wajib kosong", async () => {
      const form = new FormData();
      // name kosong, category_id=0, price=0 — semua invalid
      form.append("name", "");
      form.append("category_id", "0");
      form.append("price", "0");
      form.append("stock", "-1");

      const res = await api.post("/api/product/", form, {
        headers: {
          ...form.getHeaders(),
          Authorization: `Bearer ${authToken}`,
        },
      });

      expect([400]).toContain(res.status);
      expect(res.data.success).toBe(false);
    });

    it("harus berhasil membuat produk dengan data valid", async () => {
      // Perlu category_id yang valid, kita get dulu dari endpoint category
      const catRes = await api.get("/api/category/");
      const categories = catRes.data?.data || [];

      if (categories.length === 0) {
        console.warn("Tidak ada kategori, skip test create product");
        return;
      }

      const categoryId = categories[0].id;
      const form = new FormData();
      form.append("name", `Produk Test ${timestamp}`);
      form.append("description", "Deskripsi produk test");
      form.append("category_id", String(categoryId));
      form.append("price", "75000");
      form.append("stock", "20");

      const res = await api.post("/api/product/", form, {
        headers: {
          ...form.getHeaders(),
          Authorization: `Bearer ${authToken}`,
        },
      });

      expect(res.status).toBe(201);
      expect(res.data.success).toBe(true);
      expect(res.data.data).toHaveProperty("id");
    });
  });
});
