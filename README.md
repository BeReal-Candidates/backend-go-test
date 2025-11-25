# test-backend
## BeReal technical test for backend interview

This test is based off a real feature in our backend

It should not take you more than 2 hours

Please do **not** fork your solution to a public repository and follow this guide to share your code, you have to do it **BEFORE** you code: [HOW_TO_SHARE_MY_CODE.md](./doc/HOW_TO_SHARE_MY_CODE.md)


### The discovery feature

The Discovery timeline is a crucial BeReal feature, this is where you see all the new public content of the app

The clients (mobile applications) use the GRPC service `DiscoveryService.GetPosts` to get the new public posts

### The discovery endpoint spec

If a cursor is provided, the client wants the PAGE_SIZE posts before (older than) the cursor sorted by descending order of arrival (a scroll down in the app)

If no cursor is provided, the client wants the last PAGE_SIZE posts sorted by descending order of arrival (a pull-to-refresh in the app)

Note that:
- Deleted posts must not be returned.
- If the cursor points to a deleted post, the service must continue from the next undeleted one (i.e the next live post older than the cursor).

#### An example:

Let's say the timeline is empty, and `PAGE_SIZE = 3`.

Then post 5 is added, post 6 is added, …, post 11 is added (we add them all in the beginning for the sake of example).

The first `GetPosts` call with no cursor specified must yield the posts 11, 10, 9 and a cursor.

Using this cursor, the second call must yield the posts 8, 7, 6 and a cursor.

Using this cursor, the third call must yield the post 5 and a cursor.

Using this cursor, the fourth call must yield no posts and no cursor.

### Your goal

The feature is deliberately unfinished and it contains bugs. Your goal is to:
1. Find and fix the bugs that we planted in the code.
2. Finish the `discovery.Service` by implement the `GetPosts` and  `DeletePost` functions properly.

### Challenges to keep in mind

- Clients can have started scrolling from any post (they don't always ask for the last `PAGE_SIZE` posts)
- We don't want you to add a database engine or any go library, you can only use what Golang and this current repository provide.
- Keep the complexity low for all 3 operations (Add/Get/Delete): over a million of posts are discoverable.
- Your code should be production-ready.
