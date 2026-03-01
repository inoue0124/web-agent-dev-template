import { expect, test } from "@playwright/test";

test.describe("Home Page", () => {
  test("should display the page title", async ({ page }) => {
    await page.goto("/");
    await expect(
      page.getByRole("heading", { name: "Web Agent Dev Template" })
    ).toBeVisible();
  });

  test("should display the subtitle", async ({ page }) => {
    await page.goto("/");
    await expect(
      page.getByText("Next.js + Go/Gin で始める AI エージェント開発")
    ).toBeVisible();
  });

  test("should display stats cards", async ({ page }) => {
    await page.goto("/");
    await expect(page.getByText("Total Items")).toBeVisible();
    await expect(page.getByText("Active")).toBeVisible();
    await expect(page.getByText("Pending")).toBeVisible();
    await expect(page.getByText("Completed")).toBeVisible();
  });

  test("should display items table", async ({ page }) => {
    await page.goto("/");
    await expect(
      page.getByRole("columnheader", { name: "ID" })
    ).toBeVisible();
    await expect(
      page.getByRole("columnheader", { name: "Name" })
    ).toBeVisible();
    await expect(
      page.getByRole("columnheader", { name: "Status" })
    ).toBeVisible();
    await expect(page.getByText("Data Pipeline Setup")).toBeVisible();
  });

  test("should have New Item button", async ({ page }) => {
    await page.goto("/");
    await expect(
      page.getByRole("button", { name: "New Item" })
    ).toBeVisible();
  });
});
