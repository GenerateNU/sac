import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { Editor } from "@tiptap/react";
import { Link } from "lucide-react";
import { useCallback } from "react";
import { z } from "zod";
import { useForm, Controller } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";

interface HyperLinkProps {
  editor: Editor | null;
}

const linkSchema = z.object({
  link: z.string().url("Invalid URL"),
});

type LinkData = z.infer<typeof linkSchema>;

export const HyperlinkButton: React.FC<HyperLinkProps> = ({ editor }) => {
  if (!editor) {
    return null;
  }

  const {
    control,
    handleSubmit,
    formState: { errors },
  } = useForm<LinkData>({
    resolver: zodResolver(linkSchema),
    defaultValues: {
      link: editor.getAttributes("link").href || "",
    },
  });

  const setLinkInEditor = useCallback(
    (data: LinkData) => {
      const url = data.link;
      if (!url) {
        editor.chain().focus().extendMarkRange("link").unsetLink().run();
        return;
      }
      editor.chain().focus().extendMarkRange("link").setLink({ href: url }).run();
    },
    [editor]
  );

  const onSubmit = handleSubmit((data) => {
    setLinkInEditor(data);
  });

  return (
    <Popover>
      <PopoverTrigger asChild>
        <Button
          variant="link"
          className={editor.isActive("link") ? "bg-slate-400 rounded-md p-1.5" : "p-1.5"}
          size="link"
        >
          <Link size={20} />
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-76">
        <form onSubmit={onSubmit} className="flex flex-row items-center space-x-5">
          <div className="flex flex-row space-x-2">
            <div className="flex flex-col items-top">
              <Controller
                name="link"
                control={control}
                render={({ field }) => (
                  <Input
                    id="link"
                    placeholder="Enter a link"
                    {...field}
                    onKeyDown={(e) => {
                      if (e.key === "Enter") {
                        onSubmit();
                      }
                    }}
                  />
                )}
              />
              {errors.link && (
                <p className="text-sm pt-1 text-red-400">{errors.link.message}</p>
              )}
            </div>
            <Button type="submit">Submit</Button>
              </div>
            </form>
      </PopoverContent>
    </Popover>
  )
}
