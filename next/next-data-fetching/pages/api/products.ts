import type { NextApiRequest, NextApiResponse } from 'next'

import { getProducts } from "../../lib/products";

type Data = {
    id: number,
    title: string,
}

async function handler(req: NextApiRequest, res:NextApiResponse<Data>) {
    const products = await getProducts();
    res.status(200).json(products);
}

export default handler;