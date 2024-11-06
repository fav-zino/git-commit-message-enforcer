
---

# Git Commit Message Enforcer

A simple script to enforce standardized Git commit messages across your project. This ensures consistent commit history and helps teams maintain clean, readable project logs.

## Features

- **Commit Tag Enforcement**: Ensures commit messages start with one of the predefined tags:
  - `fix:` - Bug fix
  - `feat:` - New feature
  - `BREAKING CHANGE:` - Breaking API changes
  - `docs:` - Documentation updates
  - `test:` - Test additions or updates
  - `chore:` - Routine tasks or build system updates
  - `refactor:` - Code restructuring without behavior change
  - `build:` - Build system or dependency changes
  - `style:` - Code formatting or style adjustments
- **Validation**: Prevents commits with invalid tags, ensuring a consistent format across your project.

## Installation

1. **Download the Script**  
   Select the version of the commit message enforcer script for your operating system from the `build` folder:
   
   - **Linux**: `build/linux`
   - **MacOS**: `build/macOS`
   - **Windows**: `build/windows`

2. **Create the Hooks Folder**  
   Navigate to your Git installation directory and create a folder named `hooks` (if it doesn’t already exist).

3. **Copy the Script to the Hooks Folder**  
   Move the downloaded script to the newly created `hooks` folder and rename it as `commit-msg`.

4. **Make the Script Executable (Linux/MacOS)**  
   Run the following command to make the script executable:
   ```bash
   chmod +x ~/Git/hooks/commit-msg
   ```

5. **Apply Globally (Optional)**  
   To enforce this hook across all repositories on your system:

   Example for windows
   ```powershell
   git config --global core.hookspath "C:/Program Files/Git/hooks"
   ```
   Example for linux/MacOS
   ```bash
   git config --global core.hookspath ~/Git/hooks
   ```

6. **Reinitialize Existing Repositories**  
   For existing repositories to recognize the global hook path, run:
   ```bash
   git init
   ```
   This will not affect the repository contents or history.

7. **Verify the Hook**  
   To test the hook, make a commit. If the message format is incorrect, an error will display with guidance on the correct tags.

## Usage

Write your commit messages using one of the following formats:

```
fix: Correct a bug in the login system
feat: Add user profile feature
BREAKING CHANGE: Update authentication system API
```

If the commit message format is incorrect, you will receive an error prompt with the allowed tags.

## Example

```bash
git commit -m "feat: Add payment gateway integration"
```

**Incorrectly formatted message output:**

```
[MESSAGE FORMAT ERROR] Your commit message is not formatted correctly.
Please start your commit message with one of the following tags:
- fix:      A bug fix
- feat:     A new feature
- BREAKING CHANGE: Introduces breaking API changes
- docs:     Documentation updates
- test:     Adding or modifying tests
- perf:     Improving performance
- chore:    Routine tasks, build process, or auxiliary tools
- revert:   Revert a previous commit
- hotfix:   An urgent fix
- refactor: Code restructuring without changing functionality
- build:    Changes that affect the build system or dependencies
- style:    Code style updates (formatting, etc.)
- release:    for a release version
```


---

