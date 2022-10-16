MAIN_PATH = main.go
CWMP_PATH = internal/cwmp/main/main.go
WINDOWS_TARGET = emu-server.exe
LINUX_TARGET = rec_qq
TEMPLATE_PATH = simulation-admin/dist
INSTALL_FILE = install
ADMIN_PATH = simulation-admin

.PHONY:run
run:
	$(MAKE) download
	@echo "开始启动程序！"
	go run $(MAIN_PATH)


.PHONY:templates
templates:
	@cd $(ADMIN_PATH) && rm -rf dist
	@rm -rf templates
	@cd $(ADMIN_PATH) && yarn
ifeq ($(type), release)
	@cd $(ADMIN_PATH) && yarn build-pro
else
	@cd $(ADMIN_PATH) && yarn build-dev
endif
	@cp -rf $(TEMPLATE_PATH) templates

.PHONY:install
install:
ifneq ($(INSTALL_FILE), $(wildcard $(INSTALL_FILE)))
	@echo "install文件不存在，需重新创建！"
	@mkdir install
endif
	@rm -rf $(INSTALL_FILE)
	@mkdir install
	$(MAKE) templates type=$(type)
	@mv -f templates install/templates
	@cp -rf config install/config
	@cp -rf config.json install/config.json
	@rm -rf install/$(LINUX_TARGET).*

.PHONY:installWindows
installWindows:
	$(MAKE) install
	$(MAKE) buildWindows
	@mv $(WINDOWS_TARGET) install/$(WINDOWS_TARGET)

.PHONY:cwmp
cwmp:
	$(MAKE) download
	@echo "执行cwmp仿真器"
	go run $(CWMP_PATH)

.PHONY:buildWindows
buildWindows:
	$(MAKE) download
	@echo "打包windows二进制文件"
	EGO_ENABLED=0 go build -o $(WINDOWS_TARGET) $(MAIN_PATH)

.PHONY:buildLinux
buildLinux:
	$(MAKE) download
	@echo "打包linux二进制文件"
	EGO_ENABLED=0 go build -o $(LINUX_TARGET) $(MAIN_PATH)

.PHONY:download
download:
	@echo "下载依赖包......"
	go mod download
	go mod vendor

.PHONY:clean
clean:
	@echo "清理上一次打包的文件！"
	@rm -rf install*
	@rm -rf templates