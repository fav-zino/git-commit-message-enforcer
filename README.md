# Git Commit Message Enforcer

A simple script to enforce standardized Git commit messages across your project. This ensures consistent commit history and helps teams maintain clean, readable project logs.

## Features

- Ensures commit messages start with one of the predefined tags:
  - `fix:`: A bug fix
  - `feat:`: A new feature
  - `BREAKING CHANGE:`: Introduces breaking API changes
  - `docs:`: Documentation changes
  - `test:`: Adding or modifying tests
  - `chore:`: Routine tasks, build process, or auxiliary tools
  - `refactor:`: Code refactoring without changing functionality
  - `build:`: Changes that affect the build system or external dependencies
  - `style:`: Code style updates (formatting, etc.)
- Rejects invalid commit messages, ensuring a consistent format for your project.

## Installation

1. **Download the Script**

   Download the appropriate version of the commit message enforcer script for your operating system from the `build` folder:
   
   - **Linux**: `build/linux/commit-msg`
   - **MacOS**: `build/macOS/commit-msg`
   - **Windows**: `build/windows/commit-msg.exe`

2. **Copy the Script to Git Hooks**

   Copy the script into your local repository's `.git/hooks` folder:

   ```bash
   # For Linux
   cp build/linux/commit-msg .git/hooks/commit-msg
   chmod +x .git/hooks/commit-msg

   # For MacOS
   cp build/macOS/commit-msg .git/hooks/commit-msg
   chmod +x .git/hooks/commit-msg

   # For Windows
   copy build\windows\commit-msg.exe .git\hooks\commit-msg.exe
   ```

3. **Make the Script Executable (Linux/MacOS)**

   After copying the script, make sure it is executable:

   ```bash
   chmod +x .git/hooks/commit-msg
   ```

4. **Verify the Hook**

   You can verify that the hook is working by attempting to make a commit. If the commit message doesn't follow the required format, the script will reject it and display an error message with instructions.

## Usage

1. Write your commit messages following one of the allowed formats:

   ```
   fix: Correct a bug in the login system
   feat: Add user profile feature
   BREAKING CHANGE: Update authentication system API
   ```

2. If your message is incorrectly formatted, you'll receive a clear error message indicating the violation and the available tags.

## Example

```
git commit -m "feat: Add a new payment gateway"
```

If the message is incorrectly formatted, you will see:

```
[MESSAGE FORMAT ERROR] Your commit message is not formatted correctly.
Please start your commit message with one of the following tags:
- fix:      A bug fix
- feat:     A new feature
- BREAKING CHANGE: Introduces breaking API changes
- docs:     Documentation changes
- test:     Adding or modifying tests
- chore:    Routine tasks, build process, or auxiliary tools
- refactor: Code refactoring without changing functionality
- build:    Changes that affect the build system or external dependencies
- style:    Code style updates (formatting, etc.)
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

### Breakdown:

- **Title & Description**: Clearly explains what the project does and why it's useful.
- **Features**: Lists what the script enforces in terms of commit message formats.
- **Installation**: Provides detailed instructions on downloading, copying, and setting up the hook, including OS-specific instructions.
- **Usage**: Shows users how to use the hook and how to format commit messages.
- **Example**: Gives an example of a correct commit message and the error output when a message is incorrect.
- **License**: Standard licensing section.
