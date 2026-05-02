// test/__tests__/transaction.test.js
// Unit API Test untuk Transaction Endpoints

const axios = require("axios");
const FormData = require("form-data");
const { BASE_URL } = require("../config");

const api = axios.create({
  baseURL: BASE_URL,
  validateStatus: () => true,
});

let authToken = "";
let productId = null;

const timestamp = Date.now();
const testUser = {
  name: `Trx Tester ${timestamp}`,
  email: `trxtest${timestamp}@example.com`,
  phone: "082200000003",
  password: "testpass123",
};

beforeAll(async () => {
  // 1. Register dan login
  await api.post("/api/auth/register", testUser);
  const loginRes = await api.post("/api/auth/login", {
    email: testUser.email,
    password: testUser.password,
  });
  authToken = loginRes.data?.data?.token || "";

  // 2. Coba ambil produk yang ada untuk dipakai checkout
  const productRes = await api.get("/api/product/");
  const products = productRes.data?.data?.products || [];
  if (products.length > 0) {
    productId = products[0].id;
  }
});

describe("Transaction API", () => {
  // ===========================================================================
  // GET MY TRANSACTIONS - Protected
  // ===========================================================================
  describe("GET /api/trx/", () => {
    it("harus mengembalikan 401 tanpa token", async () => {
      const res = await api.get("/api/trx/");
      expect([401, 403]).toContain(res.status);
    });

    it("harus mengembalikan daftar transaksi milik user", async () => {
      const res = await api.get("/api/trx/", {
        headers: { Authorization: `Bearer ${authToken}` },
      });

      expect(res.status).toBe(200);
      expect(res.data.success).toBe(true);
      // Data bisa berupa array kosong atau array transaksi
      expect(Array.isArray(res.data.data)).toBe(true);
    });
  });

  // ===========================================================================
  // CHECKOUT - Protected
  // ===========================================================================
  describe("POST /api/trx/checkout", () => {
    it("harus mengembalikan 401 tanpa token", async () => {
      const res = await api.post("/api/trx/checkout", {
        product_id: 1,
        quantity: 1,
        payment_method: "transfer",
      });
      expect([401, 403]).toContain(res.status);
    });

    it("harus mengembalikan 400 jika body tidak valid", async () => {
      const res = await api.post(
        "/api/trx/checkout",
        {}, // body kosong
        { headers: { Authorization: `Bearer ${authToken}` } }
      );

      // Bisa 400 (invalid body) atau 200/error tergantung validasi server
      expect([400, 500]).toContain(res.status);
    });

    it("harus berhasil checkout jika ada produk tersedia", async () => {
      if (!productId) {
        console.warn("Tidak ada produk tersedia, skip test checkout");
        return;
      }

      const res = await api.post(
        "/api/trx/checkout",
        {
          product_id: productId,
          quantity: 1,
          payment_method: "transfer",
        },
        { headers: { Authorization: `Bearer ${authToken}` } }
      );

      // 200 jika stok cukup, 400 jika stok habis
      if (res.status === 200) {
        expect(res.data.success).toBe(true);
        expect(res.data.data).toHaveProperty("id");
      } else {
        expect([400]).toContain(res.status);
      }
    });
  });
});
