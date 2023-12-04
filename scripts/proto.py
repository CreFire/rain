import os
import subprocess

def compile_proto_files(proto_path, output_path, protoc_path, protoc_gen_go_path):
    for root, _, files in os.walk(proto_path):
        for file in files:
            if file.endswith(".proto"):
                proto_file = os.path.join(root, file)
                output_dir = os.path.join(output_path, os.path.relpath(root, proto_path))
                os.makedirs(output_dir, exist_ok=True)
                command = f"{protoc_path} --proto_path={root} --go-grpc_out={output_dir} --plugin=protoc-gen-go" \
                          f"={protoc_gen_go_path} {proto_file}"
                print(command)
                subprocess.run(command, shell=True)

if __name__ == "__main__":
    proto_path = os.path.abspath("../internal/server/proto")
    output_path = os.path.abspath("../internal/server/proto")
    protoc_path = "E:\work\go\src\\rain\scripts\protoc.exe"  # 你的 protoc.exe 路径
    protoc_gen_go_path = "E:\work\go\src\\rain\scripts\protoc-gen-go.exe"  # 你的 protoc-gen-go.exe 路径
    compile_proto_files(proto_path, output_path, protoc_path, protoc_gen_go_path)
