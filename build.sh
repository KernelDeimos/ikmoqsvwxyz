for os in windows linux darwin; do
	[[ $os = 'windows' ]] && exe='.exe' || exe=''
	for arch in amd64 386; do
		echo 'ikmoqsvwxyz_'$os'_'$arch$exe
		GOARCH=$arch GOOS=$os go build -o 'ikmoqsvwxyz_'$os'_'$arch$exe
	done
done
