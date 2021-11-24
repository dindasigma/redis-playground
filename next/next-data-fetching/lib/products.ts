import { FetchJson } from "./api";

const API_URL = 'http://localhost:1337';

export async function getProducts() {
  const products = await FetchJson(`${API_URL}/products`);
  return products.map(stripProducts);
}

function stripProducts(product: any) {
  return {
      id: product.id,
      title: product.title
  }
}