const completionSpec: Fig.Spec = {
  name: "todo",
  description: "A TODO manager just for me",
  subcommands: [
    {
      name: "list",
      description: "List TODOs",
    },
    {
      name: "add",
      description: "Add a new TODO",
      args: [
        {
          name: "title",
          description: "The title of a new TODO",
        },
        {
          name: "position",
          description: "The position of a new TODO",
          isOptional: true,
        },
      ],
      options: [
        {
          name: ["--parent", "-p"],
          description: "The parent ID of a new TODO",
          args: {
            name: "parent id",
          },
        },
        {
          name: ["--open", "-o"],
          description: "Open a new TODO file",
        }
      ]
    },
    {
      name: "open",
      description: "Open a TODO in editor",
      args: {
        name: "id",
      },
    },
    {
      name: "move",
      description: "Move a TODO",
      args: [
        {
          name: "id",
        },
        {
          name: "position",
        },
      ],
      options: [
        {
          name: ["--parent", "-p"],
          description: "The new parent ID of a moved TODO",
          args: {
            name: "id",
          },
        },
      ],
    },
    {
      name: "delete",
      description: "Delete a TODO",
      args: {
        name: "id",
        isVariadic: true,
      },
    },
    {
      name: "done",
      description: "Mark a TODO as done",
      args: {
        name: "id",
        isVariadic: true,
      },
    },
    {
      name: "undone",
      description: "Mark a TODO as undone",
      args: {
        name: "id",
        isVariadic: true,
      },
    },
    {
      name: "wait",
      description: "Mark a TODO as waiting",
      args: {
        name: "id",
        isVariadic: true,
      },
    },
    {
      name: "archive",
      description: "Archive all done TODOs",
    },
  ],
  options: [
    {
      name: ["--help", "-h"],
      description: "Show help message",
    },
    {
      name: ["--version", "-v"],
      description: "Show version",
    },
  ],
};
export default completionSpec;
