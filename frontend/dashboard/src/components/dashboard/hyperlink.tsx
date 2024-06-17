import { Button } from "../ui/button";
import { Input } from "../ui/input";
import { Popover, PopoverContent, PopoverTrigger } from "../ui/popover";
import { Editor } from "@tiptap/react";
import { Link } from "lucide-react";
import { useEffect } from "react";
import { z } from "zod";
import { useForm, Controller } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";

interface HyperLinkProps {
  editor: Editor | null;
  disabled: boolean;
}

const linkSchema = z.object({
  link: z.string().url("Invalid URL").optional(),
});

type LinkData = z.infer<typeof linkSchema>;

export const HyperlinkButton: React.FC<HyperLinkProps> = ({ editor, disabled }) => {
  const {
    control,
    handleSubmit,
    setValue,
    formState: { errors },
  } = useForm<LinkData>({
    resolver: zodResolver(linkSchema),
    defaultValues: {
      link: editor?.getAttributes("link").href || "", // Use optional chaining here
    },
  });

  useEffect(() => {
    if (!editor) return;

    const updateLink = () => {
      const currentLink = editor.getAttributes("link").href;
      setValue("link", currentLink || "");
    };

    editor.on("transaction", updateLink);
    return () => {
      editor.off("transaction", updateLink);
    };
  }, [editor, setValue]);

  const setLinkInEditor = (data: LinkData) => {
    const url = data.link;
    if (!url) {
      editor?.chain().focus().unsetLink().run();
      return;
    }
    editor?.chain().focus().extendMarkRange("link").setLink({ href: url }).run();
  };

  const onSubmit = handleSubmit((data) => {
    setLinkInEditor(data);
  });

  return (
    <Popover>
      <PopoverTrigger asChild>
        <Button
          variant="link"
          className={editor?.isActive("link") && editor?.isEditable ? "bg-slate-300 rounded-md p-1.5" : "p-1.5"}
          size="icon"
          disabled={disabled}
        >
          <Link />
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
  );
};
