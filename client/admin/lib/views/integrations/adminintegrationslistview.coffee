kd                          = require 'kd'
KDView                      = kd.View
remote                      = require('app/remote').getInstance()
KDCustomHTMLView            = kd.CustomHTMLView
KDCustomScrollView          = kd.CustomScrollView
KDListViewController        = kd.ListViewController
AdminIntegrationItemView    = require './adminintegrationitemview'
AdminIntegrationSetupView   = require './adminintegrationsetupview'
AdminIntegrationDetailsView = require './adminintegrationdetailsview'

module.exports = class AdminIntegrationsListView extends KDView

  constructor: (options = {}, data) ->

    options.cssClass = 'all-integrations'

    super options, data

    @integrationType = @getOption 'integrationType'

    @createListController()
    @fetchIntegrations()

    @addSubView @subContentView = new KDCustomScrollView

  createListController: ->

    @listController       = new KDListViewController
      viewOptions         :
        wrapper           : yes
        itemClass         : AdminIntegrationItemView
      useCustomScrollView : yes
      noItemFoundWidget   : new KDCustomHTMLView
        cssClass          : 'hidden no-item-found'
        partial           : 'No integration available.'
      startWithLazyLoader : yes
      lazyLoaderOptions   :
        spinnerOptions    :
          size            : width: 28

    @addSubView @listController.getView()


  fetchIntegrations: ->

    kd.singletons.socialapi.integrations.list (err, data) =>

      return @handleNoItem err  if err

      if @integrationType is 'configured' # dummy code.
        data = data.first

      for item in data
        item.integrationType = @integrationType
        listItem = @listController.addItem item
        @registerListItem listItem

      @listController.lazyLoader.hide()


  registerListItem: (item) ->

    item.on 'IntegrationConfigureRequested', @bound 'showIntegrationDetails'

    item.on 'IntegrationGroupsFetched', (data) =>
      setupView = new AdminIntegrationSetupView {}, data
      setupView.once 'KDObjectWillBeDestroyed', @bound 'showList'
      setupView.once 'NewIntegrationAdded',     @bound 'showIntegrationDetails'

      @showView setupView


  showView: (view) ->

    @listController.getView().hide()
    @subContentView.wrapper.destroySubViews()
    @subContentView.wrapper.addSubView view


  showList: ->

    @subContentView.wrapper.destroySubViews()
    @listController.getView().show()


  showIntegrationDetails: (data) ->

    @showView new AdminIntegrationDetailsView {}, data


  handleNoItem: (err) ->

    kd.warn err
    @listController.lazyLoader.hide()
    @listController.noItemView.show()