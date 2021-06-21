module.exports = {
  preset: 'ts-jest',
  transform: {
    '^.+\\.ts$': 'ts-jest',
    '^.+\\.(ts|js)$': 'babel-jest',
    '^.+\\.svelte$': ['svelte-jester', { preprocess: true }],
  },
  moduleFileExtensions: ['ts', 'js', 'svelte'],
  testMatch: ['**/?(*.)spec.ts'],
  setupFilesAfterEnv: ['@testing-library/jest-dom/extend-expect'],
};
