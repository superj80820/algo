using System;
using System.Collections.Generic;
using System.Linq;
namespace playground
{
    public class EmojiTrieNodeV2
    {
        public Dictionary<int, EmojiTrieNode> childrenMap;
        public Emoji emoji;

        public EmojiTrieNodeV2()
        {
            childrenMap = new();
            emoji = null;
        }
    }

    public class EmojiTrieV2
    {
        private readonly EmojiTrieNode root;
        public EmojiTrieV2()
        {
            root = new EmojiTrieNode();
        }

        public bool Insert(string id, Emoji emoji)
        {
            string[] recipe = id.Split('-');
            int[] map = new int[recipe.Length];
            for (int i = 0; i < recipe.Length; i++)
            {
                if (int.TryParse(recipe[i], out int utf32))
                {
                    map[i] = utf32;
                }
                else
                {
                    return false;
                }
            }

            EmojiTrieNode node = root;
            foreach (int key in map)
            {
                if (!node.childrenMap.ContainsKey(key))
                {
                    node.childrenMap.Add(key, new EmojiTrieNode());
                }
                node = node.childrenMap[key];
            }
            node.emoji = emoji;
            return true;
        }

        public (string, Emoji) Match(string id, int start = 0)
        {
            List<string> recipe = id.Split('-').ToList();

            Func<int, (string, Emoji)> closure = null;
            closure = start =>
            {
                int[] map = new int[recipe.Count];
                for (int i = start; i < recipe.Count; i++)
                {
                    if (int.TryParse(recipe[i], out int utf32))
                    {
                        map[i] = utf32;
                    }
                    else if (i == start) // invalid head
                    {
                        return closure(i + 1);
                    }
                    else
                    {
                        break;
                    }
                }

                EmojiTrieNode node = root;
                int j = start;
                for (; j < map.Length; j++)
                {
                    if (node.childrenMap.TryGetValue(map[j], out var nextNode))
                    {
                        node = nextNode;
                    }
                    else if (j == start)
                    {
                        return closure(j + 1);
                    }
                    else break;
                }
                if (j > start)
                {
                    string emojiId = string.Join("-", recipe.GetRange(start, j - start));
                    return (emojiId, node.emoji);
                }
                else
                {
                    return (string.Empty, null);
                }
            };

            return closure(0);
        }
    }

}





