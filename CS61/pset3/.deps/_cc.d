DEP_CC:=cc  -I. -std=gnu2x -m64 -mno-mmx -mno-sse -mno-sse2 -mno-sse3 -mno-3dnow -ffreestanding -fno-pic -fno-stack-protector -Wall -W -Wshadow -Wno-format -Wno-unused-parameter  -MD -MF .deps/.d -MP -O2 _  -Os --gc-sections -z max-page-size=0x1000 -z noexecstack -static -nostdlib -m elf_x86_64 --no-warn-rwx-segments
DEP_PREFER_GCC:=
