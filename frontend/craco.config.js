module.exports = {
    style: {
      postcss: {
        plugins: [
          require('tailwindcss'),
          require('autoprefixer'),
        ],
      },
    },
    webpack: {
      configure: {
        optimization: {
          runtimeChunk: false,
          splitChunks: {
            chunks(chunk) {
              return false
            },
          },
        },
      },
    },
  } 