import { Category } from '@/types/category';

import { tags } from './tags';

export const categories: Category[] = [];

for (let i = 1; i <= 10; i++) {
    const category: Category = {
        id: i.toString(),
        created_at: new Date(),
        updated_at: new Date(),
        name: `Tag${i}`,
        tags: tags.slice(0, Math.floor(Math.random() * tags.length))
    };
    categories.push(category);
}
