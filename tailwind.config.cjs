/** tailwind.config.cjs */
module.exports = {
  // make sure we inherit everything the theme already provides
  presets: [require("./themes/hugoplate/tailwind.config.cjs")], // adjust path if theme lives elsewhere
  theme: {
    extend: {
      spacing: {
        // add 29 ↔︎ 7.25 rem once and use w-29 / h-29 everywhere
        29: "7.25rem",
      },
    },
  },
};
