import { test, expect } from '@playwright/test';

test.describe('ads', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/');
  });

  test('has title ads-ui', async ({ page }) => {
    await expect(page).toHaveTitle(/ads-ui/);
  });
});
