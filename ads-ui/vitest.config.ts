import { defineConfig } from 'vitest/config';

export default defineConfig({
  test: {
    coverage: {
      provider: 'v8',
      reporter: ['html', 'lcov', 'text'],
      reportsDirectory: './coverage',
    },
    exclude: ['src/**/*.gen.ts'],
    include: ['src/**/*.spec.ts'],
  },
});
