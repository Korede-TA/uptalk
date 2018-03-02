# uptalk
Encrypted messaging protocol based on Upspin (upspin.io)

**WORK IN PROGRESS**

I figured that the namespacing and access models that upspin is based on might be an optimal basis for building a minimal encrpyted chat/messaging protocol.

Uptalk will assume a new directory under a user's Upspin FUSE mount point
The general layout of said directory would be something like this:

```
uptalk/
+--my_chats/
|  +--<chat_hash>/
|  |  +--Access  # Access file containing names of members of a chat
|  |  +--messages/
|  |  |  +--<msg_timestamp>.pb 
|  |  |  |  ...
|  |  |  media/
|  |  |  +--<media_file>
|  |  |  |  ...
|  |  ...
|  other_chats/
|  |  <chat_owner>_<chat_hash>
|  |  ...
|  invites/
|  +--Access  # Access file that allows writes (facilitated by uptalk) from everyone, 
|  +--<invite_timestamp>.pb
|  |  ...
```

Each chat will reside solely in the Uptalk directory of the initiating user, and will not be duplicated into other directories.
Instead, on accepting an invitation from a chat owner, the chat will be entered into the `other_chats` directory as a blank file with the chat owner and unique hash as the filename.

The `chat_hash` is generated via an MD5 hash of a string of format `<chat_owner>-<chat_members...>` i.e. the chat owner's Upsin ID followed by the IDs of each of the chat members sorted alphabetically.
This should allow for all chat hashes to be unique regardless of owner.

The .pb files will be backed by the proto message types described in [uptalk.proto](./uptalk.proto). 
The `Message` message holds the text of the message, the author of the message as well as a timestamp value. Timestamps in the Uptalk context are encoded as milliseconds since the epoch.
The `User` message only holds ab uptalk ID or "username".
The `Chat` message holds a name for the chat, the chat's owner, the chat members, invitees as well as the messages.
The `Invites` type encodes an invite to a chat owned by another user.

Given that this project is a WIP, these definitions are subject to change.
Specifically, some of the constructs that are already captured in the directory structure might be removed from the proto definition and vice-versa.

Protocol Buffers were a good choice to serialize this data because of the fast reads and writes as well as the generated code that makes it easy to interface with the data.
Also, in the event of needing to pass this data across the wire, protobufs leave me with less worry about with regards to size and speed especially when compared with JSON or XML. 
Additionally, Protobufs have built in versioning which will allow modifications to these underlying units of the protocol more painless.


