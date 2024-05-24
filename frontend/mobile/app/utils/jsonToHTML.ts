export type JsonContent = {
    type: string;
    attrs?: Record<string, any>;
    content?: JsonContent[];
    marks?: { type: string; attrs?: Record<string, any> }[];
    text?: string;
  };
  
  export const parseJsonToHtml = (json: JsonContent[]): string => {
    const parseNode = (node: JsonContent): string => {
      if (node.type === 'text') {
        let text = node.text || '';
        if (node.marks) {
          node.marks.forEach(mark => {
            switch (mark.type) {
              case 'bold':
                text = `<strong>${text}</strong>`;
                break;
              case 'italic':
                text = `<em>${text}</em>`;
                break;
              case 'strike':
                text = `<s>${text}</s>`;
                break;
              case 'underline':
                text = `<u>${text}</u>`;
                break;
              case 'link':
                const href = mark.attrs?.href || '#';
                text = `<a href="${href}" target="${mark.attrs?.target || '_self'}" rel="${mark.attrs?.rel || ''}">${text}</a>`;
                break;
            }
          });
        }
        return text;
      }
  
      let tag = node.type;
      if (node.type === 'heading' && node.attrs?.level) {
        tag = `h${node.attrs.level}`;
      } else if (node.type === 'paragraph') {
        tag = 'p';
      } else if (node.type === 'bulletList') {
        tag = 'ul';
      } else if (node.type === 'orderedList') {
        tag = 'ol';
      } else if (node.type === 'listItem') {
        tag = 'li';
      }
  
      const attrs = node.attrs
        ? ' ' +
          Object.entries(node.attrs)
            .map(([key, value]) => `${key}="${value}"`)
            .join(' ')
        : '';
  
      const content = node.content ? node.content.map(parseNode).join('') : '';
  
      return `<${tag}${attrs}>${content}</${tag}>`;
    };
  
    return json.map(parseNode).join('');
  };