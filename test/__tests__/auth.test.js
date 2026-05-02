// test/__tests__/auth.test.js
// Unit API Test untuk Auth Endpoints

const axios = require("axios");
const { BASE_URL } = require("../config");

const api = axios.create({
  baseURL: BASE_URL,
  validateStatus: () => true, // Jangan throw error untuk status apapun
});

// Data unik per test run agar tidak konflik di database
const timestamp = Date.now();
const testUser = {
  name: `Test User ${timestamp}`,
  email: `testuser${timestamp}@example.com`,
  phone: "081234567890",
  password: "password123",
};

let authToken = "";

describe("Auth API", () => {
  // ===========================================================================
  // REGISTER
  // ===========================================================================
  describe("POST /api/auth/register", () => {
    it("harus berhasil register user baru", async () => {
      const res = await api.post("/api/auth/register", testUser);

      expect(res.status).toBe(201);
      expect(res.data.success).toBe(true);
      expect(res.data.message).toMatch(/registered/i);
      expect(res.data.data).toHaveProperty("user_id");
      expect(res.data.data).toHaveProperty("email", testUser.email);
    });

    it("harus mengembalikan 400 jika field kosong", async () => {
      const res = await api.post("/api/auth/register", {
        name: "",
        email: "",
        phone: "",
        password: "",
      });

      expect(res.status).toBe(400);
      expect(res.data.success).toBe(false);
    });

    it("harus mengembalikan error jika email sudah terdaftar", async () => {
      const res = await api.post("/api/auth/register", testUser); // duplikat

      expect([400, 409]).toContain(res.status);
      expect(res.data.success).toBe(false);
    });
  });

  // ===========================================================================
  // LOGIN
  // ===========================================================================
  describe("POST /api/auth/login", () => {
    it("harus berhasil login dengan kredensial yang benar", async () => {
      const res = await api.post("/api/auth/login", {
        email: testUser.email,
        password: testUser.password,
      });

      expect(res.status).toBe(200);
      expect(res.data.success).toBe(true);
      expect(res.data.data).toHaveProperty("token");
      expect(typeof res.data.data.token).toBe("string");
      expect(res.data.data.token.length).toBeGreaterThan(0);

      // Simpan token untuk test lain
      authToken = res.data.data.token;
    });

    it("harus mengembalikan 401 dengan password salah", async () => {
      const res = await api.post("/api/auth/login", {
        email: testUser.email,
        password: "wrongpassword",
      });

      expect(res.status).toBe(401);
      expect(res.data.success).toBe(false);
    });

    it("harus mengembalikan 401 dengan email tidak terdaftar", async () => {
      const res = await api.post("/api/auth/login", {
        email: "tidak_ada@example.com",
        password: "password123",
      });

      expect(res.status).toBe(401);
      expect(res.data.success).toBe(false);
    });
  });
});

// Export token untuk test lain
module.exports = { getToken: () => authToken, testUser };
