import { Tag } from "@/types/tag";

export const tags: Tag[] = [];

for (let i = 1; i <= 10; i++) {
    const tag: Tag = {
        id: i.toString(),
        created_at: new Date(),
        updated_at: new Date(),
        name: `Tag${i}`,
    };
    tags.push(tag);
}

