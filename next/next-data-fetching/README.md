This is a [Next.js](https://nextjs.org/) project bootstrapped with [`create-next-app`](https://github.com/vercel/next.js/tree/canary/packages/create-next-app).

## Getting Started

### Development mode

```bash
yarn dev
```

### Production mode

```bash
yarn build

yarn start
```

### Data Fetching

#### [Example 1](https://github.com/dindasigma/my-playground/blob/master/next/next-data-fetching/pages/example-1.tsx) : Fetch data from variable

#### [Example 2](https://github.com/dindasigma/my-playground/blob/master/next/next-data-fetching/pages/example-2.tsx) : Fetch data on the server side (in getStaticProps).

This way the data is fetched at build time, and included in the statically generated page, it means the page loads very quickly and it also easily indexable by search engines.

#### [Example 3a](https://github.com/dindasigma/my-playground/blob/master/next/next-data-fetching/pages/example-3a.tsx) : Fetch data on the client side (in useEffect) directly from an external API

Data will be fetched on client, updated at every request.

#### [Example 3b](https://github.com/dindasigma/my-playground/blob/master/next/next-data-fetching/pages/example-3b.tsx) : Fetch data on the client side (in useEffect) from an internal API route

Data will be fetched on client, updated at every request.

#### [Example 4](https://github.com/dindasigma/my-playground/blob/master/next/next-data-fetching/pages/example-4.tsx) : Fetch data on the server side (in getStaticProps) but with Incremental Static Regeneration

Similar to Example 2 but with the addition in "revalidate" option which result Incremental Static Regeneration. This means we get all the benefits of a statically generated page, but the page will be regenerated periodically. So if the data changes in the backend, the page will also be updated.

#### [Example 5](https://github.com/dindasigma/my-playground/blob/master/next/next-data-fetching/pages/example-5.tsx) : Fetch data on the server side (in getServerSideProps)

Data will be fetched on the server side at runtime, on every request. Means the server generated HTML page dynamically on every request.

### Choosing Data Fetching Strategy

![Choosing Data Fetching Strategy](http://url/to/img.png)
Source:[https://www.udemy.com/course/nextjs-by-example/learn/lecture/27221326#overview](https://github.com/dindasigma/my-playground/blob/master/next/next-data-fetching/public/data-fetching.png)
